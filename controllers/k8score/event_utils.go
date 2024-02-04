/*
Copyright (C) 2022-2024 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package k8score

import (
	"encoding/json"
	"time"

	corev1 "k8s.io/api/core/v1"

	intctrlutil "github.com/apecloud/kubeblocks/pkg/controllerutil"
)

// IsOvertimeEvent checks whether the duration of warning event reaches the threshold.
func IsOvertimeEvent(event *corev1.Event, timeout time.Duration) bool {
	if event.Series != nil {
		return event.Series.LastObservedTime.After(event.EventTime.Add(timeout))
	}
	// Note: LastTimestamp/FirstTimestamp/Count/Source of event are deprecated in k8s v1.25
	return event.LastTimestamp.After(event.FirstTimestamp.Add(timeout))
}

// ParseProbeEventMessage parses probe event message.
func ParseProbeEventMessage(reqCtx intctrlutil.RequestCtx, event *corev1.Event) *ProbeMessage {
	message := &ProbeMessage{}
	err := json.Unmarshal([]byte(event.Message), message)
	if err != nil {
		// not role related message, ignore it
		reqCtx.Log.Info("not role message", "message", event.Message, "error", err)
		return nil
	}
	return message
}
