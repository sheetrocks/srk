fetch('../data').then(res => res.json())
.then(json => {

  // Converting from SheetRocks JSON format to appropriate JSON
  // format for d3
  let graph = {
    "nodes": [],
    "links": [],
  };
  
  let nodes = json[0].Array;
  let links = json[1].Array;
  let obj;
  
  for (var i in nodes) {
    obj = {
      "id": nodes[i][0].Text,
      "group": nodes[i][1].Number
    };
    graph.nodes.push(obj);  
  }
  for (var i in links) {
    obj = {
      "source": links[i][0].Text,
      "target": links[i][1].Text,
      "value": links[i][2].Number
    }
    graph.links.push(obj);
  }
  console.log(JSON.stringify);

  var svg = d3.select("svg"),
    width = +svg.attr("width"),
    height = +svg.attr("height");

  var color = d3.scaleOrdinal(d3.schemeCategory20);

  var simulation = d3.forceSimulation()
      .force("x", d3.forceX([width/2]).strength(0.015))
      .force("y", d3.forceY([height/2]).strength(0.015))
      .force("link", d3.forceLink().id(function(d) { return d.id; }))
      .force("charge", d3.forceManyBody())
      .force("center", d3.forceCenter(width / 2, height / 2 + 25));
      //                                                       ^
      // Due to the nature of the specific data, a 25 px offset is used
      // to make the full graph appear without having to have a tighter
      // pull on the nodes

    var link = svg.append("g")
        .attr("class", "links")
      .selectAll("line")
      .data(graph.links)
      .enter().append("line")
        .attr("stroke-width", function(d) { return Math.sqrt(d.value); });

    var node = svg.append("g")
        .attr("class", "nodes")
      .selectAll("g")
      .data(graph.nodes)
      .enter().append("g")
      
    var circles = node.append("circle")
        .attr("r", 5)
        .attr("fill", function(d) { return color(d.group); })
        .call(d3.drag()
            .on("start", dragstarted)
            .on("drag", dragged)
            .on("end", dragended));

    var lables = node.append("text")
        .text(function(d) {
          return d.id;
        })
        .attr('x', 6)
        .attr('y', 3);

    node.append("title")
        .text(function(d) { return d.id; });

    simulation
        .nodes(graph.nodes)
        .on("tick", ticked);

    simulation.force("link")
        .links(graph.links);

    function ticked() {
      link
          .attr("x1", function(d) { return d.source.x; })
          .attr("y1", function(d) { return d.source.y; })
          .attr("x2", function(d) { return d.target.x; })
          .attr("y2", function(d) { return d.target.y; });

      node
          .attr("transform", function(d) {
            return "translate(" + d.x + "," + d.y + ")";
          })
    }


  function dragstarted(d) {
    if (!d3.event.active) simulation.alphaTarget(0.3).restart();
    d.fx = d.x;
    d.fy = d.y;
  }

  function dragged(d) {
    d.fx = d3.event.x;
    d.fy = d3.event.y;
  }

  function dragended(d) {
    if (!d3.event.active) simulation.alphaTarget(0);
    d.fx = null;
    d.fy = null;
  }

})