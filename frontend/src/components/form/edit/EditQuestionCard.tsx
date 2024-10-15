import { QuestionReturn } from "../../../api/questionService"
import { QuestionType } from "../../../types/enums"
import EditMuliAnswerCard from "./EditMultiAnswerCard"
import EditOpenAnswerCard from "./EditOpenAnswerCard"
import EditSingleAnswerForm from "./EditSingleAnswerForm"

const EditQuestionCard = (props: QuestionReturn) => {
  return (
    <>
    {
        props.question_type === QuestionType.QUESTION_TYPE_MCQ && EditSingleAnswerForm(props)
    }
    {
        props.question_type === QuestionType.QUESTION_TYPE_MULTI_ANSWER_MCQ && EditMuliAnswerCard(props)
    }
    {
        props.question_type === QuestionType.QUESTION_TYPE_OPEN_ENDED && EditOpenAnswerCard(props)
    }
    </>
  )
}

export default EditQuestionCard