import Roads from "src/views/roads";

const RoadsPage = () => <Roads />;

RoadsPage.acl = {
  action: "read",
  permission: "roads",
};
export default RoadsPage;
