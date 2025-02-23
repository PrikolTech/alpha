import { computed, onMounted, onUnmounted, ref } from "vue";

export const useWindowSize = function () {
  const windowWidth = ref(window.innerWidth);
  const windowHeight = ref(window.innerHeight);

  const isMobile = computed(() => {
    return windowWidth.value < 1440;
  });

  onMounted(() => {
    window.addEventListener("resize", () => {
      windowWidth.value = window.innerWidth;
      windowHeight.value = window.innerHeight;
    });
  });
  onUnmounted(() => {
    window.removeEventListener("resize", () => {});
  });

  return { windowWidth, windowHeight, isMobile };
};
