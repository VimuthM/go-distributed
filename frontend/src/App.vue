<script setup>
import Canvas from './components/Canvas.vue'
import Input from './components/Input.vue'

let canvasRef = null;
let inputs = {
  algo: 'Chang Roberts',
  numNodes: 8,
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

<!-- <script>
export default {
  mounted() {
    console.log("Hello World!");
    renderCanvas(inputs.algo, inputs.numNodes);
  },
};
</script> -->

<template>
  <v-app>
    <v-card style="height: 100%;">
      <div style="position: relative;">
        <Canvas ref="canvasRef" style="height: 100%;" :algo="inputs.algo" :numNodes="inputs.numNodes" />
          <v-card style="padding: 20px; position: absolute; top: 0; left: 0; width: 350px; border: 1px solid #25BB4D;">
            <Input @inputs-updated="handleInputsUpdate" />
          </v-card>
      </div>
    </v-card>
  </v-app>
</template>
