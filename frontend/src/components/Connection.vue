<template>
    <div id="app">
        <button v-on:click="sendMessage('hello')">Send Message</button>
    </div>
</template>

<script>
export default {
    data: function() {
        return {
            connection: null
        }
    },
    methods: {
        sendMessage: function(message) {
            console.log("Hello")
            console.log(this.connection);
            this.connection.send(message);
        }
    },
    created: function() {
        console.log("Starting connection to WebSocket Server")
        this.connection = new WebSocket("ws://localhost:9090/ws")

        this.connection.onmessage = function(event) {
            console.log(event);
        }

        this.connection.onopen = function(event) {
            console.log(event)
            console.log("Successfully connected to the echo websocket server...")
        }
    }
}
</script>
