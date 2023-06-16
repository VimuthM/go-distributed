
<template>
<div class="greetings">
	<div id="graph"></div>
</div>
</template>

<script>

import ForceGraph from 'force-graph';

const nodes = [];
const links = [];
const nodeMap = {};
const linkMap = {};

function generateRingGraph(n, l) {
	for (let i = 0; i < n.length; i++) {
		const node = {
			id: n[i],
			name: n[i],
			val: 5,
			leader: false,
			drop: false,
		};
		nodes.push(node);
		nodeMap[node.id] = i;

		const link = {
			source: l[i][0],
			target: l[i][1],
			value: 1
		}
		links.push(link);
		linkMap[link.source] = i;
	}

	return {
		nodes,
		links,
	};
}

let graph = null;

function renderGraph(n, l) {
	console.log("Rendering");
	const numNodes = n.length;
	const g = generateRingGraph(n, l);
	const { nodes, links } = g;
	console.log(nodes, links);
	const elem = document.getElementById("graph");

	graph = ForceGraph()(elem);
	graph.graphData(g);
	graph.backgroundColor("#D0EAD8");
	graph.nodeColor(node => node.leader
		? "#10C1D8" 
		: node.drop
			? "#F9C80E"
			: "#F99597");
	graph.linkWidth(4);
	graph.zoom(1/numNodes + 4, 500)

	graph
		.linkDirectionalParticleWidth(10)
		.linkDirectionalParticleColor(link => "#02a392")
		.linkDirectionalParticleSpeed(0.015);
}

function forward(sender, receiver) {
	console.log("Forwarding: ", sender, receiver);
	graph.emitParticle(links[linkMap[sender]]);
}

function drop(id) {
	console.log("Dropping: ", id);
	nodes[nodeMap[id]].drop = true;

	setTimeout(() => {
		nodes[nodeMap[id]].drop = false;
	}, 500);
}

function leader(id) {
	console.log("Leader: ", id);
	nodes[nodeMap[id]].leader = true;
}

export default {
	methods: {
		renderGraph: renderGraph,
		forward: forward,
		drop: drop,
		leader: leader,
	}
}

</script>
