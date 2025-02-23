<template>
  <a-layout style="overflow: hidden">
    <a-layout
      :style="{
        transition: 'margin .2s',
      }"
    >
      <a-layout-content style="min-height: 100vh">
        <router-view />
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts"></script>

<style scoped lang="scss"></style>
