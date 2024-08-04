<script setup lang="ts">

const items = [
  {title: 'Home', route: '/'},
  {title: 'About', route: '/about'}
];

function click(route: string) {
  console.log(route);
}

</script>

<template>
  <v-app>
    <v-container>
      <v-app-bar :elevation="2">
        <template v-slot:prepend>
          <v-menu>
            <template v-slot:activator="{ props }">
              <v-btn color="primary" v-bind="props">
                <span class="material-symbols-outlined">
                  menu
                </span>
              </v-btn>
            </template>
            <v-list>
              <v-list-item v-for="(item, index) in items" :key="index" :value="index">
                <v-list-item-title @click="click(item.route)">{{item.title}}</v-list-item-title>
              </v-list-item>
            </v-list>
          </v-menu>
        </template>

        <v-app-bar-title>Application Bar</v-app-bar-title>

        <template v-slot:append>
          <v-btn icon="mdi-heart"></v-btn>

          <v-btn icon="mdi-magnify"></v-btn>

          <v-btn icon="mdi-dots-vertical"></v-btn>
        </template>
      </v-app-bar>
    </v-container>
  </v-app>
</template>

<style scoped>
</style>
