
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
	  color: true,
      val: 5,
    };
    nodes.push(node);

    if (i < n - 1) {
      const link = {
        source: i + 1,
        target: i + 2,
      };
      links.push(link);
    } else {
      // Connect the last node with the first node to form a ring
      const link = {
        source: n,
        target: 1,
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
    const elem = document.getElementById("graph");
    const graph = ForceGraph()(elem);
    graph.graphData(g);
	graph.backgroundColor("#D0EAD8");
	graph.nodeColor(node => node.color ? "#F99597" : "F99597");
	graph.linkWidth(4);
	graph.zoom(1/numNodes + 4, 500)
	// graph.nodeRelSize(5);


	// setInterval(() => {
	// 	const { nodes, links } = Graph.graphData();
	// 	const id = nodes.length;
	// 	const target = Math.round(Math.random() * (id - 1));
	// 	console.log(id, target);
	// 	Graph.graphData({
	// 	nodes: [...nodes, { id }],
	// 	links: [...links, { source: id, target: target }]
	// 	});
	// }, 1000);


	function removeNode(node) {
		let { nodes, links } = graph.graphData();
		links = links.filter(l => l.source !== node && l.target !== node); // Remove links attached to node
		nodes.splice(node.id, 1); // Remove node
		nodes.forEach((n, idx) => { n.id = idx; }); // Reset node ids to array index
		graph.graphData({ nodes, links });
	}
}

export default {
	methods: {
		render: render
	}
}

</script>
