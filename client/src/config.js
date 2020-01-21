export const defaultContactLink = "mailto:info@octarinsec.com?subject=Octarine%20Contact%20Request";
export const defaultWebsiteLink = "https://www.octarinesec.com";


const runtimeConfig = typeof window !== 'undefined'
    ? {
        // client
        contactLink: window.env.CONTACT_LINK || defaultContactLink,
        websiteLink: window.env.WEBSITE_LINK || defaultWebsiteLink,
    }
    : {
        // server
        contactLink: process.env.CONTACT_LINK || defaultContactLink,
        websiteLink: process.env.WEBSITE_LINK || defaultWebsiteLink,
    };

export default runtimeConfig;