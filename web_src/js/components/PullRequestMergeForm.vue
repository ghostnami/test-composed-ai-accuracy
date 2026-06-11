// hunk: @@ -3,7 +3,7 @@ import {computed, onMounted, onUnmounted, shallowRef, watch} from 'vue';
import {SvgIcon} from '../svg.ts';
import {toggleElem} from '../utils/dom.ts';

const {pageData} = window.config;

const mergeForm = pageData.pullRequestMergeForm;

// hunk: @@ -95,7 +95,6 @@ function clearMergeMessage() {

    <!-- another similar form is in pull.tmpl (manual merge)-->
    <form class="ui form form-fetch-action" v-if="showActionForm" :action="mergeForm.baseLink+'/merge'" method="post">
      <input type="hidden" name="head_commit_id" v-model="mergeForm.pullHeadCommitID">
      <input type="hidden" name="merge_when_checks_succeed" v-model="autoMergeWhenSucceed">
      <input type="hidden" name="force_merge" v-model="forceMerge">
// hunk: @@ -177,7 +176,6 @@ function clearMergeMessage() {

      <!-- the cancel auto merge button -->
      <form v-if="mergeForm.hasPendingPullRequestMerge" :action="mergeForm.baseLink+'/cancel_auto_merge'" method="post" class="tw-ml-4">
        <button class="ui button">
          {{ mergeForm.textAutoMergeCancelSchedule }}
        </button>
