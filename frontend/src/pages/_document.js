// ** React Import
import { Children } from "react";

// ** Next Import
import Document, { Html, Head, Main, NextScript } from "next/document";

class CustomDocument extends Document {
    render() {
        return (
            <Html lang="en">
                <Head>
                    <link
                        rel="preconnect"
                        href="https://fonts.googleapis.com"
                    />
                    <link rel="preconnect" href="https://fonts.gstatic.com" />
                    <link
                        rel="stylesheet"
                        href="https://fonts.googleapis.com/css2?family=Inter:wght@300;400;500;600;700&display=swap"
                    />
                    <link rel="shortcut icon" href="/images/favicon.png" />
                </Head>
                <body>
                    <Main />
                    <NextScript />
                </body>
            </Html>
        );
    }
}
CustomDocument.getInitialProps = async (ctx) => {
    const originalRenderPage = ctx.renderPage;

    ctx.renderPage = () =>
        originalRenderPage({
            enhanceApp: (App) => (props) =>
            (
                <App {...props} />
            ),
        });
    const initialProps = await Document.getInitialProps(ctx);

    return {
        ...initialProps,
        styles: [...Children.toArray(initialProps.styles)],
    };
};

export default CustomDocument;
