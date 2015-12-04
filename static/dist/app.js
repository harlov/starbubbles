
 conn = new WebSocket("ws://"+location.host+"/game_ws?name=" + name);
                conn.onopen = function() {
                    console.log("Connected");
                };
                conn.onclose = function(evt) {
                    console.log("Connection closed");
                };
                conn.onmessage = function(evt) {
                    console.log('message: ',evt.data);
                }