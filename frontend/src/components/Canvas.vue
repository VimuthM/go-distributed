
<template>
<div class="greetings">
	<div id="graph"></div>
</div>
</template>

<script>

import ForceGraph from 'force-graph';

function generateRingGraph(n) {
	const nodes = [];
	const links = [];
	for (let i = 0; i < n; i++) {
		const node = {
			id: i + 1,
			name: i + 1,
			val: 5,
			leader: false,
			receives: 0,
		};
		nodes.push(node);

		if (i < n - 1) {
			const link = {
				source: i + 1,
				target: i + 2,
				value: 1
			};
			links.push(link);
		} else {
			const link = {
				source: n,
				target: 1,
				value: 1
			};
			links.push(link);
		}
	}

	return {
		nodes,
		links,
	};
}

function render(algo, numNodes) {
	console.log("Rendering");
	console.log(algo, numNodes);
	const g = generateRingGraph(numNodes);
	const { nodes, links } = g;
	const elem = document.getElementById("graph");
	const graph = ForceGraph()(elem);
	graph.graphData(g);
	graph.backgroundColor("#D0EAD8");
	graph.nodeColor(node => node.leader ? "#10C1D8" : "#F99597");
	graph.linkWidth(4);
	graph.zoom(1/numNodes + 4, 500)

	graph
		.linkDirectionalParticleWidth(10)
		.linkDirectionalParticleColor(link => "#02a392")
		.linkDirectionalParticleSpeed(0.015);

	const emit = setInterval(() => {
		console.log("Emitting");
		const linkId = Math.floor(Math.random() * links.length);
		const recv = links[linkId].target;
		console.log(nodes, recv)
		recv.receives += 1;
		recv.leader = recv.receives > 4;
		graph.emitParticle(links[linkId]);

		if (nodes.reduce((count, node) => count + (node.leader ? 1 : 0), 0) === numNodes) {
			clearInterval(emit);
		}
	}, 1000);
}

export default {
	methods: {
		render: render,
	}
}

</script>
