const RADIUS          = 15,
      MARGIN          = 15,
      RECT_WIDTH      = 130,
      RECT_HEIGHT     = 48,
      PADDING         = 130,
      NORMAL_FONTSIZE = "1em",
      SMALL_FONTSIZE  = "0.7em";

const TYPE_PARENTS  = "parents",
      TYPE_CHILDREN = "children",
      TYPE_SELF     = "self",
      TYPE_DISABLED = "disabled";

const LABELS = {
  "GET":    "get",
  "GETALL": "list",
  "POST":   "create",
  "PUT":    "update",
  "DELETE": "delete",
}

const drawRelations = (data, nbParents, nbChildren) => {

  resizeSVG($('#relations').width(), Math.max(nbParents, nbChildren) * (RECT_HEIGHT + MARGIN) + 2* MARGIN);

  let width = $('svg').width(),
      height = $('svg').height();

  let svg = d3.select("svg")

  let legends = [
    {"x": 3, "y": 5,  "r": 3, "text": "Undocumented objects", "style":TYPE_DISABLED},
    {"x": 3, "y": 20, "r": 3, "text": "Parent objects",       "style":TYPE_PARENTS},
    {"x": 3, "y": 35, "r": 3, "text": "Current object",       "style":TYPE_SELF},
    {"x": 3, "y": 50, "r": 3, "text": "Child objects",        "style":TYPE_CHILDREN},
  ]
  var legend = svg.selectAll(".legend")
    .data(legends)
    .enter().append("g")
    .attr("class", "legend")
    .attr("class", (d) => { return d.style; })
    .attr("transform", (d, i) => {
      return "translate(" + d.x + "," + d.y + ")";
    })

  legend.append("circle")
    .attr("r", (d) => { return d.r; });

  legend.append("text")
  .text((d) => { return d.text; })
  .attr("font-size", SMALL_FONTSIZE)
  .attr("x", (d, i) => { return 8 })
  .attr("y", (d, i) => { return 3.5 });

  svg.selectAll("line")
   .data(data)
   .enter()
   .append("line")
   .attr("class", "line")
   .attr("class", (d) => { return style(d); })
   .attr("x1", (d, i) => { return getCX(d, i); })
   .attr("y1", (d, i) => { return getCY(d, i); })
   .attr("x2", (d, i) => { return getCX({type: TYPE_SELF}, 1); })
   .attr("y2", (d, i) => { return getCY({type: TYPE_SELF}, 1); })
   .attr("stroke-width", 2);

  var node = svg.selectAll(".node")
    .data(data)
    .enter().append("g")
      .attr("class", "node")
      .attr("class", (d) => { return style(d); })
      .attr("transform", (d, i) => {
        return "translate(" + getCX(d, i) + "," + getCY(d, i) + ")";
      })
      .on("click", (d) => {
        if (!d.exposed || d.type === TYPE_SELF)
          return;
        window.location.href = d.restName + ".html";
        d3.event.stopPropagation();
      })
      .style("cursor", (d,i) => {
        if (!d.exposed || d.type === TYPE_SELF)
          return "default";
        return "pointer";
      });


  node.append("rect")
    .attr("class", "box")
    .attr("x", -RECT_WIDTH / 2)
    .attr("y", -RECT_HEIGHT / 2)
    .attr("height", RECT_HEIGHT)
    .attr("width", RECT_WIDTH)
    .attr("rx", 3)
    .attr("ry", 3);


  node.append("text")
    .text((d) => { return d.instanceName; })
    .attr("class", "title")
    .style("font-size", function(d) { return adjustedSize(RECT_WIDTH, this.getComputedTextLength(), 1); })
    .attr("text-anchor", "middle");


  node.append("text")
    .text((d) => { return actions(d); })
    .style("font-size", function(d) { return adjustedSize(RECT_WIDTH, this.getComputedTextLength(), 0.8); })
    .attr("class", "actions")
    .attr("text-anchor", "middle")
    .attr("x", (d, i) => { return lineCX(d, i); })
    .attr("y", (d, i) => { return lineCY(d, i); });
}


const getCX = (d, i) => {
  switch (d.type) {
    case TYPE_PARENTS:
      return PADDING + RADIUS + MARGIN;

    case TYPE_SELF:
      return $('svg').width() / 2;

    case TYPE_CHILDREN:
      return $('svg').width() - RADIUS - MARGIN - PADDING;

    default:
      throw new Exception("Unknown type " + d.type);
  }
}

const getCY = (d, i) => {
  switch (d.type) {
    case TYPE_PARENTS:
      return (i + 1) * (RADIUS + MARGIN) * 2;

    case TYPE_SELF:
      return (RADIUS + MARGIN) * 2;

    case TYPE_CHILDREN:
      if (i == 0)
        i++;

      return i * (RADIUS + MARGIN) * 2;

    default:
      throw new Exception("Unknown type " + d.type);
  }
}

const lineCX = (d, i) => {
    return 0;
}

const lineCY = (d, i) => {
    return MARGIN;
}

const actions = (d) => {
  let results = d.actions.map((action) => {
    if (action === "GET" && d.type !== TYPE_SELF)
      action = "GETALL";

    return LABELS[action];
  });

  return results.join(", ");
}

const resizeSVG = (width, height) => {
  $('svg').attr('viewBox', "0 0 " + width + " " + height);
}

const style = (d) => {
  if (!d.exposed)
    return "ui " + TYPE_DISABLED

  return d.type;
}

const adjustedSize = (availableSpace, computedTextLength, maxValue) => {
  return  Math.min(availableSpace / (computedTextLength + MARGIN), maxValue) + "em";
}
