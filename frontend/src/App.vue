<script setup>
import Canvas from './components/Canvas.vue'
import Input from './components/Input.vue'
import Connection from './components/Connection.vue'

let connRef = null;
let canvasRef = null;

let inputs = {
  algo: 'chang_roberts',
  numNodes: 8,
}

function handleInputsUpdate(newInputs) {
  console.log("inputs updated: ", newInputs);
  inputs.algo = newInputs.algo;
  inputs.numNodes = newInputs.numNodes;
  if (connRef) {
    console.log("Sending message", inputs.algo, inputs.numNodes);
    connRef.sendMessage({
      Algorithm: inputs.algo,
      Nodes: inputs.numNodes
    });
  }
  // renderCanvas(inputs.algo, inputs.numNodes);
}

function onMessage(message) {
  if ('nodes' in message) {
    console.log("Received initial message", message);
    canvasRef.renderGraph(message.nodes, message.links);
  } else {
    console.log("Received action message", message);
    if (message.action === 'forward') {
      canvasRef.forward(message.sender_id, message.receiver_id);
    } else if (message.action === 'drop') {
      canvasRef.drop(message.sender_id);
    } else if (message.action === 'leader') {
      canvasRef.leader(message.sender_id);
    }
  }
}

</script>

<template>
  <v-app>
    <v-card style="height: 100%;">
      <div style="position: relative;">
        <Canvas ref="canvasRef" style="height: 100%;" :algo="inputs.algo" :numNodes="inputs.numNodes" />
          <v-card style="padding: 20px; position: absolute; top: 0; left: 0; width: 350px; border: 1px solid #25BB4D;">
            <Input @inputs-updated="handleInputsUpdate" />
            <Connection ref="connRef" @recv="onMessage"/>
          </v-card>
      </div>
    </v-card>
  </v-app>
</template>
