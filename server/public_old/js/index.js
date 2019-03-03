fetch("//localhost:8080/login-qr")
  .then(data => data.json())
  .then(json => {
    document.querySelector("#qr").setAttribute("src", json.image);
    const status = setInterval(() => {
      fetch("//localhost:8080/login-status")
        .then(data => data.json())
        .then(json => {
          const statusSpan = document.querySelector("#status");
          statusSpan.textContent = json.ok;

          if (json.ok) {
            clearInterval(status);
            let ws = new WebSocket("ws://localhost:8080/ws");

            ws.addEventListener("message", function(event) {
              statusSpan.textContent = event.data;
            });

            ws.addEventListener("open", function() {
              ws.send("123817257");
            });
          }
        });
    }, 3000);
  });
