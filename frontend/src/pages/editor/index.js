import CodeEditorExample from '@/views/code-editor-example'

const CodeEditorExamplePage = () => <CodeEditorExample />

CodeEditorExamplePage.acl = {
    action: 'read',
    permission: 'team-members'
}
export default CodeEditorExamplePage