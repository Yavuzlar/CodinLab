import { useRouter } from 'next/router'
import LabQuestion from 'src/views/lab-question';

const LabQuestionPage = () => {
  const router = useRouter();
  const language = router.query.language;
  const questionId = router.query.questionId;
  return <LabQuestion language={language} questionId={questionId} />;
}

LabQuestionPage.acl = {
  action: "read",
  permission: "labs",
};

export default LabQuestionPage;