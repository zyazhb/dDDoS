(function () {
  var socket = new WebSocket("ws://" + window.location.host + "/websocket/ws");

  const MessageTypes = Object.freeze({
    "DETECTED": "detect type",
    "VERIFIED": "verified type",
    "RECEIVED": "received type",
    "ML": "machine learning",
    "VOTE": "vote type",
  })

  socket.onopen = function () {
    console.log("websocket open");
  }

  socket.onclose = function() {
    console.log("websocket close");
  }

  socket.onmessage = function (event) {
    var data = JSON.parse(event.data);
    console.log("receive: ", event.data);
  }

  socket.onerror = function (event) {
    console.log(event.data);
  }
})()
