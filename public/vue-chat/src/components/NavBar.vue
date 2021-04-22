<template>
  <v-card class="overflow-hidden">
    <v-app-bar
        absolute
        color="#6A76AB"
        dark
        shrink-on-scroll
        prominent
        src="https://picsum.photos/1920/1080?random"
        fade-img-on-scroll
        scroll-target="#scrolling-techniques-3"
    >
      <template v-slot:img="{ props }">
        <v-img
            v-bind="props"
            gradient="to top right, rgba(100,115,201,.7), rgba(25,32,72,.7)"
        ></v-img>
      </template>

      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-app-bar-title>Title</v-app-bar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-dots-vertical</v-icon>
      </v-btn>
    </v-app-bar>
    <v-sheet
        id="scrolling-techniques-3"
        class="overflow-y-auto"
        max-height="1000"
    >
      <v-container style="margin-top: 10rem;">
        <Chat />
      </v-container>
    </v-sheet>
  </v-card>
</template>

<script>
import Chat from './ChatBody';
export default {
  name: "NavBar",
  components: {
    Chat,
  }
}
</script>

<style scoped>

</style>