import Labs from "src/views/labs";

const LabsPage = () => <Labs />;

LabsPage.acl = {
  action: "read",
  permission: "labs",
};
export default LabsPage;
