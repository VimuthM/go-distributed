<script setup>
import Canvas from './components/Canvas.vue'
import Input from './components/Input.vue'

let canvasRef = null;
let inputs = {
  algo: null,
  numNodes: 5,
}

function handleInputsUpdate(newInputs) {
  console.log("inputs updated: ", newInputs);
  inputs.algo = newInputs.algo;
  inputs.numNodes = newInputs.numNodes;
  renderCanvas(inputs.algo, inputs.numNodes);
}

const renderCanvas = (algo, numNodes) => {
  if (canvasRef) {
    canvasRef.render(algo, numNodes);
  }
};
</script>

<template>
  <v-app>
    <v-card style="border: 2px solid black; height: 100%;">
      <v-row>
        <v-col cols="4">
          <v-card style="height: 100%; padding: 20px;">
            <Input @inputs-updated="handleInputsUpdate" />
          </v-card>
        </v-col>
        <v-col cols="8">
          <Canvas ref="canvasRef" style="height: 100%; border: 2px solid red" :algo="inputs.algo" :numNodes="inputs.numNodes"/>
        </v-col>
      </v-row>
    </v-card>
  </v-app>
</template>
