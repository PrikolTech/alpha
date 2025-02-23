export const debounce = <T extends (...args: any[]) => void>(
  callback: T,
  delay: number,
) => {
  let debounceTimeout: ReturnType<typeof setTimeout>;

  return (...args: Parameters<T>) => {
    clearTimeout(debounceTimeout);
    debounceTimeout = setTimeout(() => {
      callback(...args);
    }, delay);
  };
};
