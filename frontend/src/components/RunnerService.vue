<template>
  <button @click="call">TEST</button>
</template>

<script setup lang="ts">
import { RunnerClient } from '../jsclient/runner.client'
import { Runner, RunRequest } from '../jsclient/runner'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'

async function call() {
  let transport = new GrpcWebFetchTransport({
    baseUrl: 'http://localhost:8080'
  })

  let client = new RunnerClient(transport)

  let req: RunRequest = {
    name: 'test',
    nodes: 1
  }

  let response = await client.runAlgo(req)
  console.log(response)
}

// let response = await client.runAlgo(req)
// console.log(response)
</script>
