"use client"
import Link from "next/link";
import {FilterIcon} from "lucide-react";
import {Badge} from "~/components/ui/badge";
import {Dialog, DialogContent} from "~/components/ui/dialog";
import {useState} from "react";
import {Button} from "~/components/ui/button";


export default function HomePage() {

    const [isDialogOpen, setDialogOpen] = useState(false);

    const openDialog = () => setDialogOpen(true);
    const closeDialog = () => setDialogOpen(false);

    {/*TODO: Look to make this search better*/}
    return (
        <div className="p-4">
            <div className="pb-4 border-b">
                <div className="flex items-center justify-between">
                    <h2 className="text-2xl font-bold">Audit log</h2>
                    <div className="relative">
                        <Button
                            onClick={openDialog}
                            type="button"
                            className="flex items-center px-4 py-2 text-white bg-blue-500 rounded-md">
                            <FilterIcon className="mr-2" />
                            Filter
                            {/*<Badge className="ml-2">2</Badge>*/}
                        </Button>
                    </div>
                </div>
            </div>
            <div className="mt-4">
                <div className="overflow-x-auto">
                    <table className="min-w-full bg-white border border-gray-200">
                        <thead>
                        <tr className="text-left">
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Member</th>
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Action</th>
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Location</th>
                            <th className="px-4 py-2 font-medium text-gray-600 border-b">Date</th>
                        </tr>
                        </thead>
                        <tbody>
                        {/* Example row, replace with your data */}
                        <tr>
                            <td className="px-4 py-2 border-b">No results found</td>
                            <td className="px-4 py-2 border-b">We couldn't find any items matching your filter parameters.</td>
                            <td className="px-4 py-2 border-b"></td>
                            <td className="px-4 py-2 border-b"></td>
                        </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <Dialog open={isDialogOpen} onOpenChange={closeDialog}>
                <DialogContent>
                <form action="#" className="space-y-4">
                    <div className="flex justify-between mb-4">
                        <button type="button" className="text-blue-500" onClick={() => alert('Clear clicked')}>Clear</button>
                        <p className="font-semibold">Filters</p>
                        <button type="submit" className="text-blue-500">Apply</button>
                    </div>

                    <div className="space-y-4">
                        <div className="flex space-x-4">
                            <label className="flex items-center">
                                <input type="checkbox" className="form-checkbox" />
                                <span className="ml-2">Account member</span>
                            </label>
                            <select name="actorId" className="border border-gray-300 rounded-md p-2">
                                <option value=""></option>
                                <option value="c67f430a-a3fd-485e-8bc2-22e346e56ed8">Zac Clifton</option>
                            </select>
                        </div>

                        <div className="flex space-x-4">
                            <label className="flex items-center">
                                <input type="checkbox" className="form-checkbox" />
                                <span className="ml-2">Action</span>
                            </label>

                            <select name="action" className="border border-gray-300 rounded-md p-2">
                                <option value="account.enable">account.enable</option>
                                <option value="account.disable">account.disable</option>
                                <option value="account.sso_connected">account.sso_connected</option>
                                <option value="account.sso_disconnected">account.sso_disconnected</option>
                                <option value="account.dsync_connected">account.dsync_connected</option>
                                <option value="account.dsync_disconnected">account.dsync_disconnected</option>
                                <option value="account.dsync_user_provisioned">account.dsync_user_provisioned</option>
                                <option value="account.dsync_user_deprovisioned">account.dsync_user_deprovisioned</option>
                                <option value="account.dsync_group_user_added">account.dsync_group_user_added</option>
                                <option value="account.dsync_group_user_removed">account.dsync_group_user_removed</option>
                                <option value="account.dsync_group_deleted">account.dsync_group_deleted</option>
                                <option value="account.autojoin_updated">account.autojoin_updated</option>
                                <option value="account_invite.created">account_invite.created</option>
                                <option value="account_invite.revoked">account_invite.revoked</option>
                                <option value="account_invite.resent">account_invite.resent</option>
                                <option value="account_invite.accepted">account_invite.accepted</option>
                                <option value="seat.removed">seat.removed</option>
                                <option value="seat.role_changed">seat.role_changed</option>
                                <option value="service_token.created">service_token.created</option>
                                <option value="service_token.renamed">service_token.renamed</option>
                                <option value="service_token.deleted">service_token.deleted</option>
                                <option value="asset.created">asset.created</option>
                                <option value="asset.archived">asset.archived</option>
                                <option value="channel.created">channel.created</option>
                                <option value="channel.updated">channel.updated</option>
                                <option value="channel.archived">channel.archived</option>
                                <option value="channel.environment_settings_updated">channel.environment_settings_updated</option>
                                <option value="channel_group.created">channel_group.created</option>
                                <option value="channel_group.updated">channel_group.updated</option>
                                <option value="email_layout.created">email_layout.created</option>
                                <option value="email_layout.updated">email_layout.updated</option>
                                <option value="email_layout.duplicated">email_layout.duplicated</option>
                                <option value="email_layout.archived">email_layout.archived</option>
                                <option value="email_layout.published">email_layout.published</option>
                                <option value="email_layout.reset_to_published_version">email_layout.reset_to_published_version</option>
                                <option value="email_layout.reverted_to_target_version">email_layout.reverted_to_target_version</option>
                                <option value="api_key.created">api_key.created</option>
                                <option value="api_key.deleted">api_key.deleted</option>
                                <option value="environment.api_keys_revoked">environment.api_keys_revoked</option>
                                <option value="environment.api_key_disabled">environment.api_key_disabled</option>
                                <option value="environment.api_key_enabled">environment.api_key_enabled</option>
                                <option value="environment.merged_latest">environment.merged_latest</option>
                                <option value="environment.created">environment.created</option>
                                <option value="environment.deleted">environment.deleted</option>
                                <option value="environment.updated">environment.updated</option>
                                <option value="jwt_signing_key.generated">jwt_signing_key.generated</option>
                                <option value="translation.created">translation.created</option>
                                <option value="translation.updated">translation.updated</option>
                                <option value="translation.archived">translation.archived</option>
                                <option value="translation.reset_to_published_version">translation.reset_to_published_version</option>
                                <option value="translation.reverted_to_target_version">translation.reverted_to_target_version</option>
                                <option value="partial.created">partial.created</option>
                                <option value="partial.updated">partial.updated</option>
                                <option value="partial.archived">partial.archived</option>
                                <option value="partial.reset_to_published_version">partial.reset_to_published_version</option>
                                <option value="partial.reverted_to_target_version">partial.reverted_to_target_version</option>
                                <option value="variable.created">variable.created</option>
                                <option value="variable.updated">variable.updated</option>
                                <option value="variable.upserted">variable.upserted</option>
                                <option value="variable.removed">variable.removed</option>
                                <option value="webhook.created">webhook.created</option>
                                <option value="webhook.status_updated">webhook.status_updated</option>
                                <option value="webhook.archived">webhook.archived</option>
                                <option value="workflow.created">workflow.created</option>
                                <option value="workflow.updated">workflow.updated</option>
                                <option value="workflow.step_created">workflow.step_created</option>
                                <option value="workflow.step_cloned">workflow.step_cloned</option>
                                <option value="workflow.step_updated">workflow.step_updated</option>
                                <option value="workflow.steps_updated">workflow.steps_updated</option>
                                <option value="workflow.step_deleted">workflow.step_deleted</option>
                                <option value="workflow.steps_reordered">workflow.steps_reordered</option>
                                <option value="workflow.step_template_updated">workflow.step_template_updated</option>
                                <option value="workflow.status_updated">workflow.status_updated</option>
                                <option value="workflow.archived">workflow.archived</option>
                                <option value="workflow.cloned">workflow.cloned</option>
                                <option value="workflow.template_cloned">workflow.template_cloned</option>
                                <option value="workflow.reset_to_published_version">workflow.reset_to_published_version</option>
                                <option value="workflow.reverted_to_target_version">workflow.reverted_to_target_version</option>
                                <option value="external_user.preferences_updated">external_user.preferences_updated</option>
                                <option value="external_object.preferences_updated">external_object.preferences_updated</option>
                                <option value="integration_source_event_action_mapping.created">integration_source_event_action_mapping.created</option>
                                <option value="integration_source_event_action_mapping.updated">integration_source_event_action_mapping.updated</option>
                                <option value="integration_source_event_action_mapping.deleted">integration_source_event_action_mapping.deleted</option>
                                <option value="integration_source_event_action_mapping.status_updated">integration_source_event_action_mapping.status_updated</option>
                                <option value="integration_source_event_action_mapping.reset_to_published_version">integration_source_event_action_mapping.reset_to_published_version</option>
                                <option value="integration_source_event_action_mapping.reverted_to_target_version">integration_source_event_action_mapping.reverted_to_target_version</option>
                                <option value="integration_source.created">integration_source.created</option>
                                <option value="integration_source.deleted">integration_source.deleted</option>
                                <option value="integration_source_environment_settings.created">integration_source_environment_settings.created</option>
                                <option value="integration_source_environment_settings.deleted">integration_source_environment_settings.deleted</option>
                                <option value="integration_source_environment_settings.updated">integration_source_environment_settings.updated</option>
                            </select>
                        </div>
                    </div>
                </form>
                </DialogContent>
            </Dialog>

        </div>
    );
}
