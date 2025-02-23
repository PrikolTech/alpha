import { notification } from "ant-design-vue";

export type notificationType = "success" | "info" | "warning" | "error";

export type notificationPlacement =
  | "topLeft"
  | "topRight"
  | "bottomLeft"
  | "bottomRight";

export const createNotification = (
  type: notificationType = "info",
  title: string = type === "error"
    ? "Ошибка!"
    : type === "success"
      ? "Успешно!"
      : "",
  desc: string = "",
  placement: notificationPlacement = "topRight",
  duration?: number,
) => {
  const finalDuration =
    duration !== undefined ? duration : type === "error" ? 15 : 7;

  notification[type]({
    message: title,
    description: desc,
    placement: placement,
    duration: finalDuration,
  });
};
