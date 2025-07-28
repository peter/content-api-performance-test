#!/usr/bin/env node

import _ from 'lodash';
import axios from 'axios';
import assert from 'node:assert';

const BASE_URL = process.env.BASE_URL || 'http://localhost:8888';
const TEST_LIMIT = parseInt(process.env.TEST_LIMIT || 10_000)
const TEST_PARALLEL = parseInt(process.env.TEST_PARALLEL || 100)
const TEST_DATA_FIELD = process.env.TEST_DATA_FIELD !== 'false'
const TEST_CREATE_ONLY = process.env.TEST_CREATE_ONLY === 'true'

const LOG_LEVELS = ['DEBUG', 'INFO', 'ERROR']
const LOG_LEVEL = (process.env.LOG_LEVEL || 'DEBUG').toUpperCase()
const LOG_LEVEL_INDEX = LOG_LEVELS.indexOf(LOG_LEVEL)
if (LOG_LEVEL_INDEX === -1) {
    throw new Error(`Invalid log level: ${LOG_LEVEL}`)
}

function log(level, message, data = {}) {
    if (LOG_LEVEL_INDEX > LOG_LEVELS.indexOf(level)) {
        return
    }
    const timestamp = new Date().toISOString()
    console.log(JSON.stringify({ timestamp, level, message, ...data }))
}

function logDebug(message, data = {}) {
    log('DEBUG', message, data)
}

function logInfo(message, data = {}) {
    log('INFO', message, data)
}

function logError(message, data = {}) {
    log('ERROR', message, data)
}

function stats(values) {
    const sortedValues = _.sortBy(values)
    return {
        count: values.length,
        min: _.min(values),
        max: _.max(values),
        avg: _.mean(values),
        p90: sortedValues[Math.floor(sortedValues.length * 0.9)],
        p95: sortedValues[Math.floor(sortedValues.length * 0.95)],
        p99: sortedValues[Math.floor(sortedValues.length * 0.99)],
    }
}

function assertRecentDate(date) {
    if (!date || isNaN(date.getTime())) {
        throw new Error(`Invalid date: ${date}`)
    }
    const now = new Date()
    const diff = now.getTime() - date.getTime()
    const diffSeconds = diff / 1000
    if (diffSeconds > 10) {
        throw new Error(`Date is too old: ${date} (diff: ${diffSeconds}s)`)
    }
}

async function runTest(batchIndex, index) {
    const runId = `${batchIndex}-${index}-${Math.random().toString(36).substring(2, 15)}`
    const testStartTime = Date.now()
    const requestElapsed = { create: [], read: [], update: [], delete: [] }
    const createdAt = new Date().toISOString()
    const content = {
		title:  `Smoke Test Content ${runId}`,
		body:   `This is smoke test content number ${runId}`,
		author: "Smoke Tester",
		status: "draft",
    }
    if (TEST_DATA_FIELD) {
        content.data = {
			"run_id": runId,
			"created_at": createdAt,
		}
    }

    // CREATE
    let startTime = Date.now()
    let response = await axios.post(`${BASE_URL}/content`, content)
    requestElapsed.create.push(Date.now() - startTime)
    const id = response.data.id

    if (!TEST_CREATE_ONLY) {
        // READ
        startTime = Date.now()
        response = await axios.get(`${BASE_URL}/content/${id}`)
        requestElapsed.read.push(Date.now() - startTime)
        assert.strictEqual(response.data.id, id)
        assert.strictEqual(response.data.title, content.title)
        assert.strictEqual(response.data.body, content.body)
        assert.strictEqual(response.data.author, content.author)
        assert.strictEqual(response.data.status, content.status)
        assertRecentDate(new Date(response.data.created_at))
        if (TEST_DATA_FIELD) {
            assert.strictEqual(response.data.data.run_id, runId)
            assert.strictEqual(response.data.data.created_at, createdAt)
        }

        // UPDATE
        startTime = Date.now()
        await axios.put(`${BASE_URL}/content/${id}`, {
            ...content,
            status: "published",
            title: `${content.title} (updated)`,
        })
        requestElapsed.update.push(Date.now() - startTime)

        // READ
        startTime = Date.now()
        response = await axios.get(`${BASE_URL}/content/${id}`)
        requestElapsed.read.push(Date.now() - startTime)
        assert.strictEqual(response.data.id, id)
        assert.strictEqual(response.data.title, `${content.title} (updated)`)
        assertRecentDate(new Date(response.data.updated_at))

        // DELETE
        startTime = Date.now()
        await axios.delete(`${BASE_URL}/content/${id}`)
        requestElapsed.delete.push(Date.now() - startTime)

        // READ
        startTime = Date.now()
        await axios.get(`${BASE_URL}/content/${id}`, { validateStatus: (status) => status === 404 })
        requestElapsed.read.push(Date.now() - startTime)
    }

    return {
        testElapsed: Date.now() - testStartTime,
        requestElapsed,
        requestCount: Object.values(requestElapsed).flat().length,
    }
}

async function runBatch(batchIndex, batch) {
    const startTime = Date.now()
    logDebug(`Starting batch`, { batchIndex, batchSize: batch.length })
    const batchResults = await Promise.allSettled(batch.map(async (_, index) => {
        try {
            const result = await runTest(batchIndex, index);
            return result;
        } catch (error) {
            logError(`Error thrown running test`, { batchIndex, index, error: error.stack || error })
            throw error;
        }
    }))
    logDebug(`Finished batch`, { batchIndex, batchSize: batch.length, elapsed: Date.now() - startTime })
    return batchResults
}

async function main() {
    const startTime = Date.now()
    let results = []
    const tests = _.range(0, TEST_LIMIT)
    const batches = _.chunk(tests, TEST_PARALLEL)
    logDebug(`Starting performance test`, { TEST_LIMIT, TEST_PARALLEL, N_BATCHES: batches.length })
    for (const [batchIndex, batch] of batches.entries()) {
        const batchResults = await runBatch(batchIndex, batch)
        results = results.concat(batchResults)
    }

    const elapsedTotal = Date.now() - startTime
    const testCount = results.length
    const errorCount = results.filter(result => result.status !== 'fulfilled').length
    const successResults = results.filter(result => result.status === 'fulfilled')
    const requestCount = results.reduce((acc, result) => acc + (result.value?.requestCount || 0), 0)
    const testElapsed = successResults.map(result => _.sum(result.value.testElapsed))
    const createElapsed = successResults.map(result => result.value.requestElapsed.create).flat()
    const readElapsed = successResults.map(result => result.value.requestElapsed.read).flat()
    const updateElapsed = successResults.map(result => result.value.requestElapsed.update).flat()
    const deleteElapsed = successResults.map(result => result.value.requestElapsed.delete).flat()
    const requestElapsed = successResults.map(result => Object.values(result.value.requestElapsed).flat()).flat()
    logInfo(`Finished performance test`, {
        TEST_LIMIT,
        TEST_PARALLEL,
        N_BATCHES: batches.length,
        testCount: {
            error: errorCount,
            success: testCount - errorCount,
            total: testCount,
        },
        testElapsed: stats(testElapsed),
        createElapsed: stats(createElapsed),
        readElapsed: stats(readElapsed),
        updateElapsed: stats(updateElapsed),
        deleteElapsed: stats(deleteElapsed),
        requestElapsed: stats(requestElapsed),
        requests: {
            totalCount: requestCount,
            countPerSecond: requestCount / (elapsedTotal / 1000),
        },
        elapsedTotal,
    })
}

main()
