<script setup lang="ts">

import * as grpc from '../jsclient/runner/runner_grpc_web_pb.js';
// import { Runner, RunnerClient } from '../jsclient/runner/runner_pb_service.js';
// import { Runner } from '../jsclient/runner/runner_pb_service.js';

const client = new grpc.RunnerClient("http://localhost:8080");

const stream = client.runAlgo(new grpc.Runner.runAlgo.requestType());

stream.on('data', (response) => {
  console.log('Received response:', response);
});

stream.on('status', (status) => {
  console.log('Received status:', status);
});

stream.on('end', (status) => {
  console.log('Call ended with status:', status);
});

</script>

<template>
  <div class="greetings">
    <h1 class="green">{{ "WAH" }}</h1>
    <h3>
      You’ve successfully created a project with
      <a href="https://vitejs.dev/" target="_blank" rel="noopener">Vite</a> +
      <a href="https://vuejs.org/" target="_blank" rel="noopener">Vue 3</a>. What's next?
    </h3>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
