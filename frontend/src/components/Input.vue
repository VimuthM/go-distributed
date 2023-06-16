<template>
	<v-sheet class="mx-auto">

		<v-form ref="form">
			<v-select
				v-model="algo"
				:items="algos"
				item-title="label"
				item-value="value"
				:rules="[v => !!v || 'Algo is required']"
				label="Algo"
				required
			></v-select>

			<p class="font-weight-regular">
				&nbsp {{ numNodes  }} nodes
			</p>

			<v-slider
				v-model="numNodes"
				step="1"
				:min="1"
				:max="15"
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
				Run
				</v-btn>
			</div>
		</v-form>
	</v-sheet>
</template>

<script>
  export default {
    data() { return {
      valid: true,
      numNodes: 8,
      numNodesRules: [
        v => !!v || 'Number of nodes required',
        v => (v && Number.isInteger(Number(v)) && v > 0) || 'Number of nodes must be a positive integer',
      ],
      algo: 'chang_roberts',
      algos: [
		{ label: 'Chang Roberts', value: 'chang_roberts' },
        { label: 'Spanning Tree', value: 'spanning_tree' }
      ],
      tickLabels: Object.fromEntries(Array.from({length: 15}, (_, i) => [i + 1, i + 1]))
	}},

    methods: {
      	async validateAndRender () {
        	const { valid } = await this.$refs.form.validate()
        	if (!valid) return
			this.$emit('inputs-updated', { algo: this.algo, numNodes: this.numNodes });
      	},
    },
  }
</script>