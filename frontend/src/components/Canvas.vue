
<template>
  <div class="greetings">
    <div id="graph"></div>
    <v-btn @click="render">Render</v-btn>
  </div>
</template>

<script setup>

import ForceGraph from 'force-graph';

const g = {
    "nodes": [
        {
          "id": "id1",
          "name": "name1",
          "val": 1
        },
        {
          "id": "id2",
          "name": "name2",
          "val": 10
        }
    ],
    "links": [
        {
            "source": "id1",
            "target": "id2"
        }
    ]
}

function render() {
    console.log("rendering");
    const elem = document.getElementById("graph");
    const Graph = ForceGraph()(elem)
        // .onNodeClick(removeNode)
        .graphData(g);

	setInterval(() => {
		const { nodes, links } = Graph.graphData();
		const id = nodes.length;
		const target = Math.round(Math.random() * (id - 1));
		console.log(id, target);
		Graph.graphData({
		nodes: [...nodes, { id }],
		links: [...links, { source: id, target: target }]
		});
	}, 1000);


	function removeNode(node) {
		let { nodes, links } = Graph.graphData();
		links = links.filter(l => l.source !== node && l.target !== node); // Remove links attached to node
		nodes.splice(node.id, 1); // Remove node
		nodes.forEach((n, idx) => { n.id = idx; }); // Reset node ids to array index
		Graph.graphData({ nodes, links });
	}
}

// export default {
// 	props: ['algo', 'numNodes'],
// 	render: render
// }

</script>


<!-- <style scoped>
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
</style> -->
