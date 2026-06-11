// hunk: @@ -5,7 +5,7 @@ import {initRepoSettingsBranchesDrag} from './repo-settings-branches.ts';
import {fomanticQuery} from '../modules/fomantic/base.ts';
import {globMatch} from '../utils/glob.ts';

const {appSubUrl} = window.config;

function initRepoSettingsCollaboration() {
  // Change collaborator access mode
// hunk: @@ -56,7 +56,6 @@ function initRepoSettingsSearchTeamBox() {
    rawResponse: true,
    apiSettings: {
      url: `${appSubUrl}/org/${searchTeamBox.getAttribute('data-org-name')}/teams/-/search?q={query}`,
      onResponse(response: any) {
        const items: Array<Record<string, any>> = [];
        for (const item of response.data) {
