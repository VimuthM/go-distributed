<script>
export default {
    data: function() {
        return {
            connection: null
        }
    },
    methods: {
        sendMessage: function(message) {
            this.connection.send(JSON.stringify(message));
        }
    },
    created: function() {
        console.log("Starting connection to WebSocket Server")
        this.connection = new WebSocket("ws://localhost:9090/ws")

        this.connection.onmessage = (event) => {
            this.$emit("recv", JSON.parse(event.data));
        };

        this.connection.onopen = (event) => {
            console.log(event)
            console.log("Successfully connected to the echo websocket server...")
        }

        this.connection.onclose = (event) => {
            console.log(event)
            console.log("Connection closed")
        }
    }
}
</script>
