

export enum UserRole {
  Manager = 'manager',
  Creator = 'creator',
  Admin = 'admin',
  Readonly = 'read-only',
  Unknown = 'unknown',
}

export const EditingUserRoles = new Set([
  UserRole.Manager,
  UserRole.Admin,
  UserRole.Creator,
]);