import { useRouter } from "next/router";
import RoadDetails from "src/views/road-detail";

const RoadDetailPage = () => {
  const router = useRouter();
  const language = router.query.language;

  return <RoadDetails language={language} />;
};

RoadDetailPage.acl = {
  action: "read",
  permission: "roads",
};

export default RoadDetailPage;
