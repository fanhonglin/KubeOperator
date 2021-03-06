import {Component, OnInit} from '@angular/core';
import {SystemService} from '../system.service';
import {CommonAlertService} from '../../../layout/common-alert/common-alert.service';
import {TranslateService} from '@ngx-translate/core';
import {BaseModelComponent} from '../../../shared/class/BaseModelComponent';
import {System, SystemCreateRequest} from '../system/system';
import {AlertLevels} from '../../../layout/common-alert/alert';

@Component({
    selector: 'app-ldap',
    templateUrl: './ldap.component.html',
    styleUrls: ['./ldap.component.css']
})
export class LdapComponent extends BaseModelComponent<System> implements OnInit {

    item: System = new System();
    createItem: SystemCreateRequest = new SystemCreateRequest();

    constructor(private systemService: SystemService, private commonAlertService: CommonAlertService,
                private translateService: TranslateService) {
        super(systemService);
    }

    ngOnInit(): void {
        this.listSystemSettings();
    }


    listSystemSettings() {
        this.systemService.singleGet().subscribe(res => {
            this.item = res;
        });
    }

    onSubmit() {
        this.createItem.vars = this.item.vars;
        this.systemService.ldapCreate(this.createItem).subscribe(res => {
            this.commonAlertService.showAlert(this.translateService.instant('APP_ADD_SUCCESS'), AlertLevels.SUCCESS);
        }, error => {
            this.commonAlertService.showAlert(error.error.msg, AlertLevels.ERROR);
        });
    }

    onSync() {
        this.createItem.vars = this.item.vars;
        this.systemService.ldapSync(this.createItem).subscribe(res => {
            this.commonAlertService.showAlert(this.translateService.instant('APP_SYNC_NOTE'), AlertLevels.SUCCESS);
        }, error => {
            this.commonAlertService.showAlert(error.error.msg, AlertLevels.ERROR);
        });
    }
}
