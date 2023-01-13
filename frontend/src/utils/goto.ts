import {BrowserOpenURL} from "../../wailsjs/runtime";

function goToUrl(url: string) {
    BrowserOpenURL(url)
}

export {
    goToUrl
}
