import Sample from "@/views/sample"

const ChallengesPage = () => <Sample />

ChallengesPage.acl = {
    action: 'read',
    permission: 'sample'
}
export default ChallengesPage