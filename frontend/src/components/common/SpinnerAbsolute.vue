<template>
  <div class="example" :class="{ absolute: absolute, background: background }">
    <a-spin :size="size" />
  </div>
</template>

<script setup lang="ts">
interface Props {
  absolute?: boolean;
  background?: boolean;
  size?: "small" | "default" | "large";
}

withDefaults(defineProps<Props>(), {
  size: "default",
});
</script>

<style scoped lang="scss">
.example {
  position: fixed;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  border-radius: 4px;
  display: flex;
  justify-content: center;
  align-items: center;
  &.absolute {
    position: absolute;
  }
  &.background {
    background: rgba(0, 0, 0, 0.3);
  }
}
</style>
