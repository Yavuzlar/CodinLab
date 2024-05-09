import TeamMembers from "src/views/team-members";

const TeamMembersPage = () => <TeamMembers />;

TeamMembersPage.acl = {
  action: "read",
  permission: "team-members",
};
export default TeamMembersPage;
