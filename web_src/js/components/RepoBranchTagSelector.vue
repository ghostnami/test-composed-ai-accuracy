// hunk: @@ -28,7 +28,6 @@ export default defineComponent({
  data() {
    const shouldShowTabBranches = this.elRoot.getAttribute('data-show-tab-branches') === 'true';
    return {
      allItems: [] as ListItem[],
      selectedTab: (shouldShowTabBranches ? 'branches' : 'tags') as SelectedTab,
      searchTerm: '',
// hunk: @@ -272,7 +271,6 @@ export default defineComponent({
            {{ textCreateRefFrom.replace('%s', currentRefShortName) }}
          </div>
          <form ref="createNewRefForm" method="post" :action="createNewRefFormActionUrl">
            <input type="hidden" name="new_branch_name" :value="searchTerm">
            <input type="hidden" name="create_tag" :value="String(selectedTab === 'tags')">
            <input type="hidden" name="current_path" :value="currentTreePath">
