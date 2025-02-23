<template>
  <a-empty :image="simpleImage" :description="desc" />
</template>

<script setup lang="ts">
import { Empty } from "ant-design-vue";

const simpleImage = Empty.PRESENTED_IMAGE_SIMPLE;

defineProps<{
  desc: string;
}>();
</script>

<style scoped></style>
