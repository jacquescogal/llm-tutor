import { QuestionType } from "./enums";

export interface QuestionOption {
  optionId: number;
  optionBody: string;
}

export type QuestionProps = {
  questionType: QuestionType;
  questionId: number;
  questionBody: string;
  optionList?: QuestionOption[];
  answerOptions?: number[];
  answerOption?: number;
  answer?: string;
};
