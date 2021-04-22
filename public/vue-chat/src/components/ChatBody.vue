<template>
  <v-container>
    <v-row>
      <v-card>
        <v-card-text>

        </v-card-text>
      </v-card>
    </v-row>
    <v-row v-if="joined">
      <v-col cols="8">
        <v-text-field
            v-model="message"
            label="Message"
            filled
            @keyup.enter="send"
        ></v-text-field>
      </v-col>
      <v-col cols="4">
        <v-btn class="ma-2" @click="audioInput" :color="isSpeaking? 'error' : 'primary'">
          {{isSpeaking? 'Stop Audio' : 'Start Audio'}}
          <v-icon v-if="!isSpeaking" :src="micImage" right>mdi-microphone</v-icon>
          <v-icon v-else :src="stopImage" right> mdi-microphone-off </v-icon>
        </v-btn>
        <v-btn @click="send">
          <v-icon>mdi-chat</v-icon>
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
import SpeechToText from '../services/speech-to-text';
import micImage from '../assets/mic.svg';
import stopImage from '../assets/stop.svg';

export default {
  name: "ChatBody",
  components: {

  },
  data() {
    return {
      ws: null,

      email: null,
      message: null,
      username: null,

      participants: [],
      messages: [
          {sender: 'default name', sender_email: 'r@g.com', content: 'This is the default content'}
      ],

      joined: false,

      // speech to text
      isSpeaking: false,
      speech: '',
      speechService: {},
      micImage,
      stopImage
    };
  },
  methods: {
    send() {
      if (this.message !== "") {
        this.ws.send(
            JSON.stringify({
              email: this.email,
              username: this.username,
              message: this.message
            })
        );
        this.message = "";
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
      this.joined = true;
    },
    audioInput(){
      this.isSpeaking = true;
      this.speechService.speak().subscribe(
          result => {
            this.isSpeaking = false;
            this.message = result;
            console.log('Result', result);
          },
          (err) => {
            console.error(err);
            this.isSpeaking = false;
          },
          () => {
            this.isSpeaking = false;
          }
      );
      console.log('speechService started');
    }
  },
  mounted() {
    this.ws = new WebSocket("ws://" + window.location.host + "/ws");
    this.ws.addEventListener("message", e => {
      let msg = JSON.parse(e.data);
      console.log("Message", msg)
      let exists = false
      this.participants.map(p => {
        if (p.username === msg.username) exists = true
      })
      if (!exists) {
        this.participants.push({
          email: msg.email,
          username: msg.username
        })
      }

      this.messages.push({
        sender: msg.username,
        sender_email: msg.email,
        content: msg.message
      })
    });
  },
  created() {
    this.speechService = new SpeechToText();
  }
};
</script>

<style scoped>
</style>
