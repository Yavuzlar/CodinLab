import Sample from "src/views/sample";

const ChallengesPage = () => <Sample />;

ChallengesPage.acl = {
  action: "read",
  permission: "home",
};
export default ChallengesPage;
