import { useRouter } from 'next/router'
import LanguageLab from 'src/views/language-lab';

const LanguageLabPage = () => {
    const router = useRouter();
    const language = router.query.language;
    return <LanguageLab language={language}/>;
}
 
LanguageLabPage.acl = {
  action: "read",
  permission: "labs",
};

export default LanguageLabPage;