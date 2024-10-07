import React from 'react'
import { QuestionOption, QuestionType } from '../../../types/question'
import EditSingleAnswerForm from './EditSingleAnswerForm';
import EditMultiAnswerCard from './EditMultiAnswerCard';
import EditOpenAnswerCard from './EditOpenAnswerCard';

  

export type EditQuestionProps = {
    questionType: QuestionType;
    questionId: string;
    questionBody: string;
    optionList?: QuestionOption[];
    answerOptions?: number[];
    answerOption?: number;
    answer?: string;
}

const EditQuestionCard = (props: EditQuestionProps) => {
  return (
    <>
    {
        props.questionType === QuestionType.singleAnswer && EditSingleAnswerForm(props)
    }
    {
        props.questionType === QuestionType.multiAnswer && EditMultiAnswerCard(props)
    }
    {
        props.questionType === QuestionType.openEnded && EditOpenAnswerCard(props)
    }
    </>
  )
}

export default EditQuestionCard