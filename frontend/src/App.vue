<script setup>
import Canvas from './components/Canvas.vue'
import Input from './components/Input.vue'
</script>

<template>
  <v-app>
    <v-card style="border: 1px solid brown;">
      <v-row>
        <v-col cols="4">
          <v-card style="height: 100%; padding: 20px;">
            <Input />
          </v-card>
        </v-col>
        <v-col cols="8">
          <v-card style="height: 100%;">
            <Canvas />
          </v-card>
        </v-col>
      </v-row>
    </v-card>
  </v-app>
</template>
