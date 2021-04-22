<template>
  <v-app>
    <v-main>
      <NavBar />
      <Chat/>
    </v-main>
  </v-app>
</template>

<script>
import Chat from './components/Chat';
import NavBar from "@/components/NavBar";

export default {
  name: 'App',

  components: {
    Chat,
    NavBar,
  },

  data: () => ({
  }),
};
</script>
