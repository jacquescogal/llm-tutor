import { OrderByField } from "../types/enums";

export const sortByMap: { [key: string]: OrderByField } = {
    "ID": OrderByField.ORDER_BY_FIELD_ID,
    "Title": OrderByField.ORDER_BY_FIELD_ID,
    "Created At": OrderByField.ORDER_BY_FIELD_CREATED_TIME,
    "Updated At": OrderByField.ORDER_BY_FIELD_UPDATED_TIME,
  }