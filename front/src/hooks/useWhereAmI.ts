import { Ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useAccountStore } from '@/store';
import { getQueryStringParams, localStorageActions } from '@/common/util';

export const useWhereAmI = (): {
  whereAmI: Ref<Senarios>;
  isResourcePage: boolean;
  isBusinessPage: boolean;
  isServicePage: boolean;
  isWorkbenchPage: boolean;
  isSchemePage: boolean;
  getBusinessApiPath: () => string;
} => {
  const route = useRoute();
  const senario = computed(() => {
    if (!route) return;
    if (/^\/resource\/.+$/.test(route?.path)) return Senarios.resource;
    if (/^\/business\/.+$/.test(route.path)) return Senarios.business;
    if (/^\/service\/.+$/.test(route.path)) return Senarios.service;
    if (/^\/workbench\/.+$/.test(route.path)) return Senarios.workbench;
    if (/^\/scheme\/.+$/.test(route.path)) return Senarios.scheme;
    if (/^\/bill\/.+$/.test(route.path)) return Senarios.bill;
    if (/^\/403\/.+$/.test(route.path)) return Senarios.unauthorized;
    return Senarios.unknown;
  });

  /**
   * @returns 业务下需要拼接的 API 路径
   */
  const getBusinessApiPath = () => {
    const { bizs } = useAccountStore();

    return senario.value === Senarios.business
      ? `bizs/${bizs || getQueryStringParams('bizs') || localStorageActions.get('bizs')}/`
      : '';
  };

  return {
    whereAmI: senario,
    isResourcePage: senario.value === Senarios.resource,
    isBusinessPage: senario.value === Senarios.business,
    isServicePage: senario.value === Senarios.service,
    isWorkbenchPage: senario.value === Senarios.workbench,
    isSchemePage: senario.value === Senarios.scheme,
    getBusinessApiPath,
  };
};

export enum Senarios {
  business = 'business',
  resource = 'resource',
  service = 'service',
  workbench = 'workbench',
  scheme = 'scheme',
  bill = 'bill',
  unknown = 'unknown',
  unauthorized = 'unauthorized',
}
