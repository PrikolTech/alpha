<template>
  <a-alert :message="textDisplay" type="error" show-icon style="width: 100%">
    <template #action v-if="tryAgain">
      <a-button size="small" danger @click="$emit('clickTryAgainBtn')">
        Повторить попытку
      </a-button>
    </template>
  </a-alert>
</template>

<script setup lang="ts">
import { computed } from "vue";

interface Props {
  tryAgain?: boolean;
  text?: string;
}

const props = withDefaults(defineProps<Props>(), {
  tryAgain: false,
  text: "",
});

const textDisplay = computed(() => {
  return props.text || "Произошла ошибка при загрузке данных";
});
</script>

<style scoped></style>
