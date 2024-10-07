export interface QuestionOption {
  optionId: number;
  optionBody: string;
}

export enum QuestionType {
  singleAnswer = "singleAnswer",
  multiAnswer = "multiAnswer",
  openEnded = "openEnded",
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
