<template>
    <v-app>
<v-sheet width="700" class="mx-auto">

    <v-form ref="form">
        <v-select
            v-model="algo"
            :items="algos"
            :rules="[v => !!v || 'Algo is required']"
            label="Algo"
            required
        ></v-select>

        <div class="text-caption">Number of Nodes: {{ slider  }}</div>

        <v-slider
            v-model="slider"
            step="1"
            :min="1"
            :max="15"
            show-ticks
            tick-size="4"
            :ticks="tickLabels"
        ></v-slider>

        <div class="d-flex flex-column">
            <v-btn
            color="success"
            class="mt-4"
            block
            @click="validateAndRender"
            >
            Render
            </v-btn>
        </div>
    </v-form>
</v-sheet>
</v-app>
</template>

<script>
  export default {
    data: () => ({
      valid: true,
      numNodes: 5,
      numNodesRules: [
        v => !!v || 'Number of nodes required',
        v => (v && Number.isInteger(Number(v)) && v > 0) || 'Number of nodes must be a positive integer',
      ],
      algo: null,
      algos: [
        'Chang Roberts',
        'Spanning Tree',
      ],
      slider: 5,
      tickLabels: Object.fromEntries(Array.from({length: 15}, (_, i) => [i + 1, i + 1]))
    }),

    methods: {
      async validateAndRender () {
        const { valid } = await this.$refs.form.validate()

        if (valid) alert('Form is valid')
      },
    },
  }
</script>