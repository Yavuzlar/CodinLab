import { useRouter } from "next/router";
import LanguageRoads from "src/views/language-road";

const LanguageRoadPage = () => {
  const router = useRouter();
  const language = router.query.language;
  const pathId = router.query.pathId;
  return <LanguageRoads language={language} pathId={pathId} />;
};

LanguageRoadPage.acl = {
  action: "read",
  permission: "roads",
};

export default LanguageRoadPage;
