// Copyright 2018 Dave Evans. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package flasharray

import (
	"fmt"
	"testing"
)

const testAccVolumeName = "testAccvolume"
const testvolsnapshot = "testaccvolsnapshot"
const testvolclone = "testaccvolcone"
const testvolsize = 1024000000
const testvolresize = 2048000000
const testpgroup = "testacchostpgroup"

func TestAccVolumes(t *testing.T) {
	testAccPreChecks(t)
	c := testAccGenerateClient(t)

	c.Protectiongroups.CreateProtectiongroup(testpgroup, nil)

	t.Run("CreateVolume", testAccCreateVolume(c))
	t.Run("GetVolume", testAccGetVolume(c))
	t.Run("GetVolume_withParamSpace", testAccGetVolumeWithParamSpace(c))
	t.Run("GetVolume_withParamAction", testAccGetVolumeWithParamAction(c))
	t.Run("CreateSnapshot", testAccCreateSnapshot(c))
	t.Run("CloneVolume", testAccCloneVolume(c))
	t.Run("AddVolumeToPgroup", testAccAddVolume(c, testpgroup))
	t.Run("RemoveVolumeFromPgroup", testAccRemoveVolume(c, testpgroup))
	t.Run("ExtendVolume", testAccExtendVolume(c))
	t.Run("TruncateVolume", testAccTruncateVolume(c))
	t.Run("ListVolumes", testAccListVolumes(c))
	t.Run("ListVolumes_withParams", testAccListVolumesWithParams(c))
	t.Run("RenameVolume", testAccRenameVolume(c, "testAccVolnew"))
	c.Volumes.RenameVolume("testAccVolnew", testAccVolumeName)
	t.Run("DeleteVolume", testAccDeleteVolume(c))
	t.Run("RecoverVolume", testAccRecoverVolume(c))
	c.Volumes.DeleteVolume(testAccVolumeName)
	t.Run("EradicateVolume", testAccEradicateVolume(c))

	c.Volumes.DeleteVolume(testvolclone)
	c.Volumes.DeleteVolume(fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot))
	c.Volumes.EradicateVolume(testvolclone)
	c.Volumes.EradicateVolume(fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot))
	c.Protectiongroups.DestroyProtectiongroup(testpgroup)
	c.Protectiongroups.EradicateProtectiongroup(testpgroup)
}

func testAccCreateVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CreateVolume(testAccVolumeName, testvolsize)
		if err != nil {
			t.Fatalf("error creating volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
	}
}

func testAccGetVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.GetVolume(testAccVolumeName, nil)
		if err != nil {
			t.Fatalf("error getting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
	}
}

func testAccGetVolumeWithParamSpace(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		h, err := c.Volumes.GetVolume(testAccVolumeName, params)
		if err != nil {
			t.Fatalf("error getting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
		if h.Size != testvolsize {
			t.Fatalf("expected: %d; got %d", testvolsize, h.Size)
		}
	}
}

func testAccGetVolumeWithParamAction(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.GetVolume(testAccVolumeName, map[string]string{"action": "monitor"})
		if err != nil {
			t.Fatalf("error getting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Name)
		}
		if h.Time == "" {
			t.Fatalf("time property did not exist; got: %+v", h)
		}
	}
}

func testAccCreateSnapshot(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CreateSnapshot(testAccVolumeName, testvolsnapshot)
		if err != nil {
			t.Fatalf("error snapshotting volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot) {
			t.Fatalf("expected: %s; got %s", fmt.Sprintf("%s.%s", testAccVolumeName, testvolsnapshot), h.Name)
		}
	}
}

func testAccCloneVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		h, err := c.Volumes.CopyVolume(testvolclone, testAccVolumeName, false)
		if err != nil {
			t.Fatalf("error cloning volume %s: %s", testAccVolumeName, err)
		}

		if h.Name != testvolclone {
			t.Fatalf("expected: %s; got %s", testvolclone, h.Name)
		}
		if h.Source != testAccVolumeName {
			t.Fatalf("expected: %s; got %s", testAccVolumeName, h.Source)
		}
	}
}

func testAccAddVolume(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.AddVolume(testAccVolumeName, pgroup)
		if err != nil {
			t.Fatalf("error adding volume %s to pgroup %s: %s", testAccVolumeName, pgroup, err)
		}
	}
}

func testAccRemoveVolume(c *Client, pgroup string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RemoveVolume(testAccVolumeName, pgroup)
		if err != nil {
			t.Fatalf("error removing volume %s from pgroup %s: %s", testAccVolumeName, pgroup, err)
		}
	}
}

func testAccListVolumes(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.ListVolumes(nil)
		if err != nil {
			t.Fatalf("error listing volumes: %s", err)
		}
	}
}

func testAccListVolumesWithParams(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		params := map[string]string{"space": "true"}
		_, err := c.Volumes.ListVolumes(params)
		if err != nil {
			t.Fatalf("error listing volumes: %s", err)
		}
	}
}

func testAccExtendVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.ExtendVolume(testAccVolumeName, testvolresize)
		if err != nil {
			t.Fatalf("error extending volume %s to %d: %s", testAccVolumeName, testvolresize, err)
		}
	}
}

func testAccTruncateVolume(c *Client) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.TruncateVolume(testAccVolumeName, testvolsize)
		if err != nil {
			t.Fatalf("error truncating volume %s to %d: %s", testAccVolumeName, testvolsize, err)
		}
	}
}

func testAccRenameVolume(c *Client, name string) func(*testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RenameVolume(testAccVolumeName, name)
		if err != nil {
			t.Fatalf("error renaming volume %s to %s: %s", testAccVolumeName, name, err)
		}
	}
}

func testAccDeleteVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.DeleteVolume(testAccVolumeName)
		if err != nil {
			t.Fatalf("error deleting volume: %s", err)
		}
	}
}

func testAccRecoverVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.RecoverVolume(testAccVolumeName)
		if err != nil {
			t.Fatalf("error recovering volume %s: %s", testAccVolumeName, err)
		}
	}
}

func testAccEradicateVolume(c *Client) func(t *testing.T) {
	return func(t *testing.T) {
		_, err := c.Volumes.EradicateVolume(testAccVolumeName)
		if err != nil {
			t.Fatalf("error eradicating volume: %s", err)
		}
	}
}
