<template>
  <v-card class="overflow-hidden">
    <v-app-bar
        absolute
        color="#161B22"
        hide-on-scroll
        scroll-target="#scrolling-techniques-3"
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-app-bar-title>Socket-Chat</v-app-bar-title>

      <v-spacer></v-spacer>

      <v-btn icon>
        <v-icon>mdi-cog</v-icon>
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
      <v-container style="margin-top: 6rem;">
        <router-view />
      </v-container>
    </v-sheet>

    <v-navigation-drawer
        v-model="drawer"
        absolute
        bottom
        temporary
    >
      <v-list
          nav
          dense
      >
        <v-list-item-group
            v-model="group"
            active-class="deep-purple--text text--accent-4"
        >
          <v-list-item>
            <v-list-item-title>Chat</v-list-item-title>
          </v-list-item>

          <v-list-item>
            <v-list-item-title>TextArea</v-list-item-title>
          </v-list-item>

          <v-list-item>
            <v-list-item-title>Data Collection</v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>
  </v-card>
</template>

<script>
export default {
  name: "NavBar",
  components: {

  },
  data: () => ({
    drawer: false,
    group: null,
  }),

  watch: {
    group () {
      this.drawer = false
    },
  },
}
</script>

<style scoped>

</style>