<template>
  <v-container>
    <v-row>
      <v-card>
        <v-card-text v-html="chatContent">

        </v-card-text>
      </v-card>
    </v-row>
    <v-row v-if="joined">
      <v-col cols="8">
        <v-text-field
            v-model="newMsg"
            label="Message"
            filled
            @keyup.enter="send"
        ></v-text-field>
      </v-col>
      <v-col cols="4">
        <v-btn @click="send">
          <v-icon>chat</v-icon>
          Send
        </v-btn>
      </v-col>
    </v-row>

    <v-row v-if="!joined">
      <v-form>
        <v-container>
          <v-row>
            <v-col
                cols="12"
                sm="6"
            >
              <v-text-field
                  v-model="email"
                  label="Email"
                  filled
              ></v-text-field>
            </v-col>

            <v-col
                cols="12"
                sm="6"
            >
              <v-text-field
                  v-model="username"
                  label="Username"
                  filled
              ></v-text-field>
            </v-col>
          </v-row>
          <v-btn @click="join()">
            <v-icon>done</v-icon>
            Join
          </v-btn>
        </v-container>
      </v-form>
    </v-row>
  </v-container>
</template>

<script>
import CryptoJs from 'crypto-js'
import emojione from 'emojione'

export default {
  name: "Chat",
  data() {
    return {
      ws: null, // Our websocket
      newMsg: "", // Holds new messages to be sent to the server
      chatContent: "", // A running list of chat messages displayed on the screen
      email: null, // Email address used for grabbing an avatar
      username: null, // Our username
      joined: false, // True if email and username have been filled in
    };
  },
  methods: {
    send() {
      if (this.newMsg !== "") {
        this.ws.send(
            JSON.stringify({
              email: this.email,
              username: this.username,
              message: this.message
            })
        );
        this.newMsg = ""; // Reset newMsg
      }
    },

    join() {
      if (!this.email) {
        alert("You must enter an email");
        return;
      }
      if (!this.username) {
        alert("You must choose a username");
        return;
      }
      // this.email = ("<p>").html(this.email).text();
      // this.username = ("<p>").html(this.username).text();
      this.joined = true;
    },

    gravatarURL(email) {
      return "http://www.gravatar.com/avatar/" + CryptoJs.MD5(email);
    },
  },
  mounted() {
    let self = this;
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.addEventListener("message", function (e) {
      let msg = JSON.parse(e.data);
      self.chatContent +=
          '<div class="chip">' +
          '<img src="' +
          self.gravatarURL(msg.email) +
          '">' + // Avatar
          msg.username +
          "</div>" +
          emojione.toImage(msg.message) +
          "<br/>"; // Parse emojis

      let element = document.getElementById("chat-messages");
      element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
    });
  },
};
</script>

<style scoped>
#chat-messages {
  min-height: 10vh;
  height: 60vh;
  width: 100%;
  overflow-y: scroll;
}
</style>
